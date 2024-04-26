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
        "github.com/magicbutton/magic-mix/services/endpoints/transformer"
                    "github.com/magicbutton/magic-mix/services/models/transformermodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestTransformercreate(t *testing.T) {
                                // transformer v1
            object := transformermodel.Transformer{}
         
            result,err := transformer.TransformerCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
