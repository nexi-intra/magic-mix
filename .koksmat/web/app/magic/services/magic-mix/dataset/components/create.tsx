    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/koksmat/useservice";
    import { useState } from "react";
    import {DatasetForm} from "./form";
    
    import {DatasetItem} from "../applogic/model";
    export default function CreateDataset() {
       
        const dataset = {} as DatasetItem;
        return (
          <div>{dataset && 
          <DatasetForm dataset={dataset} editmode="create"/>}
         
          </div>
        );
    }
