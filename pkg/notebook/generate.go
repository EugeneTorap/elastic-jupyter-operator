package notebook

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tkestack/jupyter-operator/api/v1alpha1"
)

const (
	defaultImage         = "jupyter/base-notebook:python-3.8.6"
	defaultContainerName = "notebook"
	defaultPortName      = "notebook"
	defaultPort          = 8888

	LabelNotebook = "notebook"
	LabelNS       = "namespace"
)

type generator struct {
	nb *v1alpha1.JupyterNotebook
}

// newGenerator creates a new Generator.
func newGenerator(nb *v1alpha1.JupyterNotebook) (
	*generator, error) {
	if nb == nil {
		return nil, fmt.Errorf("Got nil when initializing Generator")
	}
	g := &generator{
		nb: nb,
	}

	return g, nil
}

func (g generator) DesiredDeploymentWithoutOwner() *appsv1.Deployment {
	labels := g.labels()
	selector := &metav1.LabelSelector{
		MatchLabels: labels,
	}
	d := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: g.nb.Namespace,
			Name:      g.nb.Name,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: selector,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            defaultContainerName,
							Image:           defaultImage,
							ImagePullPolicy: v1.PullIfNotPresent,
							Args: []string{
								"start-notebook.sh",
							},
							Ports: []v1.ContainerPort{
								{
									Name:          defaultPortName,
									ContainerPort: defaultPort,
									Protocol:      v1.ProtocolTCP,
								},
							},
						},
					},
				},
			},
		},
	}

	if g.nb.Spec.Gateway != nil {
		gatewayURL := fmt.Sprintf("http://%s.%s:%d",
			g.nb.Spec.Gateway.Name, g.nb.Spec.Gateway.Namespace, defaultPort)
		d.Spec.Template.Spec.Containers[0].Args = append(
			d.Spec.Template.Spec.Containers[0].Args, "--gateway-url", gatewayURL)
	}
	if g.nb.Spec.Resources != nil {
		d.Spec.Template.Spec.Containers[0].Resources = *g.nb.Spec.Resources
	}

	return d
}

func (g generator) labels() map[string]string {
	return map[string]string{
		LabelNS:       g.nb.Namespace,
		LabelNotebook: g.nb.Name,
	}
}
