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
    import {JobForm} from "./form";
    
    import {JobItem} from "../applogic/model";
    export default function CreateJob() {
       
        const job = {} as JobItem;
        return (
          <div>{job && 
          <JobForm job={job} editmode="create"/>}
         
          </div>
        );
    }
