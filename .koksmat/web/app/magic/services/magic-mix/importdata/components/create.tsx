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
    import {ImportDataForm} from "./form";
    
    import {ImportDataItem} from "../applogic/model";
    export default function CreateImportData() {
       
        const importdata = {} as ImportDataItem;
        return (
          <div>{importdata && 
          <ImportDataForm importdata={importdata} editmode="create"/>}
         
          </div>
        );
    }
