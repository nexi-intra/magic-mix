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
    import {ProcessLogForm} from "./form";
    
    import {ProcessLogItem} from "../applogic/model";
    export default function CreateProcessLog() {
       
        const processlog = {} as ProcessLogItem;
        return (
          <div>{processlog && 
          <ProcessLogForm processlog={processlog} editmode="create"/>}
         
          </div>
        );
    }
