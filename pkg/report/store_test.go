package report_test

import (
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kyverno/policy-reporter/pkg/crd/api/policyreport/v1alpha2"
	"github.com/kyverno/policy-reporter/pkg/report"
)

func Test_PolicyReportStore(t *testing.T) {
	store := report.NewPolicyReportStore()
	store.CreateSchemas()

	t.Run("Add/Get", func(t *testing.T) {
		_, ok := store.Get(preport.GetID())
		if ok == true {
			t.Fatalf("Should not be found in empty Store")
		}

		store.Add(preport)
		_, ok = store.Get(preport.GetID())
		if ok == false {
			t.Errorf("Should be found in Store after adding report to the store")
		}
	})

	t.Run("Update/Get", func(t *testing.T) {
		ureport := &v1alpha2.PolicyReport{
			ObjectMeta: v1.ObjectMeta{
				Name:              "polr-test",
				Namespace:         "test",
				CreationTimestamp: v1.Now(),
			},
			Results: make([]v1alpha2.PolicyReportResult, 0),
			Summary: v1alpha2.PolicyReportSummary{Skip: 1},
		}

		store.Add(preport)
		r, _ := store.Get(preport.GetID())
		if r.GetSummary().Skip != 0 {
			t.Errorf("Expected Summary.Skip to be 0")
		}

		store.Update(ureport)
		r2, _ := store.Get(preport.GetID())
		if r2.GetSummary().Skip != 1 {
			t.Errorf("Expected Summary.Skip to be 1 after update")
		}
	})

	t.Run("Delete/Get", func(t *testing.T) {
		_, ok := store.Get(preport.GetID())
		if ok == false {
			t.Errorf("Should be found in Store after adding report to the store")
		}

		store.Remove(preport.GetID())
		_, ok = store.Get(preport.GetID())
		if ok == true {
			t.Fatalf("Should not be found after Remove report from Store")
		}
	})

	t.Run("CleanUp", func(t *testing.T) {
		store.Add(preport)

		store.CleanUp()
		_, ok := store.Get(preport.GetID())
		if ok == true {
			t.Fatalf("Should have no results after CleanUp")
		}
	})
}
