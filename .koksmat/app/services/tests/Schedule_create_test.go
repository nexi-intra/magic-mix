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
        "github.com/magicbutton/magic-mix/services/endpoints/schedule"
                    "github.com/magicbutton/magic-mix/services/models/schedulemodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestSchedulecreate(t *testing.T) {
                                // transformer v1
            object := schedulemodel.Schedule{}
         
            result,err := schedule.ScheduleCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
