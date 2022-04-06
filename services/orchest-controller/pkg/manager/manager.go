package manager

import (
	orchestv1alpha1 "github.com/orchest/orchest/services/orchest-controller/pkg/apis/orchest/v1alpha1"
	"github.com/orchest/orchest/services/orchest-controller/pkg/reconciler/orchestcluster"
	"github.com/orchest/orchest/services/orchest-controller/pkg/utils"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Manager encapsulates creating kubernetes controller manager.
type Manager struct {
	scheme *runtime.Scheme
	config *rest.Config
}

// NewManager returns *NewManager.
func NewManager(inCluster bool) *Manager {

	scheme := runtime.NewScheme()
	orchestv1alpha1.AddToScheme(scheme)

	config := utils.GetClientConfig(inCluster)

	clientgoscheme.AddToScheme(scheme)

	return &Manager{
		scheme: scheme,
		config: config,
	}
}

// Run the operator instance.
func (m *Manager) Run() error {

	mgr, err := ctrl.NewManager(
		m.config, ctrl.Options{
			Scheme: m.scheme,
		})
	if err != nil {
		return err
	}

	reconciler := orchestcluster.NewOrchestClusterReconciler(mgr)

	if err := ctrl.NewControllerManagedBy(mgr).
		For(&orchestv1alpha1.OrchestCluster{}).
		Complete(reconciler); err != nil {
		return err
	}

	klog.Info("starting orchest manager")
	return mgr.Start(ctrl.SetupSignalHandler())

}
