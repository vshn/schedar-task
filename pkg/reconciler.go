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
	// Client Kubernetes client
	client.Client

	// Scheme for the client
	Scheme *runtime.Scheme
}

// Reconcile reconciles the CRD after some time to make sure the resource is in correct state
func (s *SimplePostgreSQLReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	//TODO implement me
	pg := v1alpha1.SimplePostgreSQL{}
	c := v1.SGPostgresConfig{}
	sg := v1.SGCluster{}

	// Deletion
	if !pg.DeletionTimestamp.IsZero() {
		return reconcile.Result{}, nil
	}

	// Creation
	if pg.Generation == 1 {

	}

	fmt.Print(sg)
	fmt.Print(c)

	return reconcile.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (s *SimplePostgreSQLReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.SimplePostgreSQL{}).
		Complete(s)
}
