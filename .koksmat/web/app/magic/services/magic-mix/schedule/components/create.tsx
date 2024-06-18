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
    import {ScheduleForm} from "./form";
    
    import {ScheduleItem} from "../applogic/model";
    export default function CreateSchedule() {
       
        const schedule = {} as ScheduleItem;
        return (
          <div>{schedule && 
          <ScheduleForm schedule={schedule} editmode="create"/>}
         
          </div>
        );
    }
