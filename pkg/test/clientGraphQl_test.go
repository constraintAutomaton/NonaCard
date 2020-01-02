package test

import (
	"testing"

	clientGraphQl "github.com/constraintAutomaton/anilist-user-analysis/pkg/clientGraphQl"
)

func TestNbOccurence(t *testing.T) {

	if value := clientGraphQl.NbOccurence("aaa", "a"); value != 3 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 3)
	}

	if value := clientGraphQl.NbOccurence("aaa", "b"); value != 0 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 0)
	}

	if value := clientGraphQl.NbOccurence("", "b"); value != 0 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 0)
	}

	if value := clientGraphQl.NbOccurence("aaa", ""); value != 0 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 0)
	}

	if value := clientGraphQl.NbOccurence("aaa", "aa"); value != 1 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 1)
	}

	if value := clientGraphQl.NbOccurence("ababace", "b"); value != 2 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 2)
	}

	if value := clientGraphQl.NbOccurence("ababace", "e"); value != 1 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 1)
	}
	if value := clientGraphQl.NbOccurence("a", "a"); value != 1 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 1)
	}

	if value := clientGraphQl.NbOccurence("a", "b"); value != 0 {
		t.Errorf("NbOccurence was incorrect, got: %d, want: %d.", value, 0)
	}

}
