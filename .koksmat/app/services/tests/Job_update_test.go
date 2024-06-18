    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/magic-mix/services/endpoints/job"
                    "github.com/magicbutton/magic-mix/services/models/jobmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestJobupdate(t *testing.T) {
                                // transformer v1
            object := jobmodel.Job{}
         
            result,err := job.JobUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
