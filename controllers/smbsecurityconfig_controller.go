/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sambaoperatorv1alpha1 "github.com/samba-in-kubernetes/samba-operator/api/v1alpha1"
)

// SmbSecurityConfigReconciler reconciles a SmbSecurityConfig object
type SmbSecurityConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//revive:disable kubebuilder directives

// +kubebuilder:rbac:groups=samba-operator.samba.org,resources=smbsecurityconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=samba-operator.samba.org,resources=smbsecurityconfigs/status,verbs=get;update;patch

//revive:enable

// Reconcile the SmbSecurityConfig resource.
func (r *SmbSecurityConfigReconciler) Reconcile(
	_ context.Context, req ctrl.Request) (ctrl.Result, error) {
	// ---
	_ = r.Log.WithValues("smbsecurityconfig", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the reconciler.
func (r *SmbSecurityConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sambaoperatorv1alpha1.SmbSecurityConfig{}).
		Complete(r)
}
