package pkg

import (
	"context"
	"fmt"
	"github.com/vshn/appcat/v4/apis/stackgres/v1"
	"github.com/vshn/schedar-task/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// SimplePostgreSQLReconciler reconciles a VSHNPostgreSQL object
type SimplePostgreSQLReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (s *SimplePostgreSQLReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	//TODO implement me
	sg := v1.SGCluster{}
	fmt.Print(sg)
	return reconcile.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (s *SimplePostgreSQLReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.SimplePostgreSQL{}).
		Complete(s)
}
