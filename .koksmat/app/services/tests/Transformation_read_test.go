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
        
        "github.com/stretchr/testify/assert"
    )
    
    func TestTransformationread(t *testing.T) {
                    
            result,err := transformation.TransformationRead(-1)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
