    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/app/koksmat/useservice";
    import { useState } from "react";
    import {TransformerForm} from "./form";
    
    import {TransformerItem} from "../applogic/model";
    export default function CreateTransformer() {
       
        const transformer = {} as TransformerItem;
        return (
          <div>{transformer && 
          <TransformerForm transformer={transformer} editmode="create"/>}
         
          </div>
        );
    }
