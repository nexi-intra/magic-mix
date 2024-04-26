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
        "github.com/magicbutton/magic-mix/services/endpoints/transformation"
                    "github.com/magicbutton/magic-mix/services/models/transformationmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestTransformationcreate(t *testing.T) {
                                // transformer v1
            object := transformationmodel.Transformation{}
         
            result,err := transformation.TransformationCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
