    
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/       
"use client";
import { z } from "zod";
// spunk
// Schedule
export interface ScheduleItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    cron : string ;
    job_id : number ;
    data : object ;

}


// Schedule
export const ScheduleSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    cron : z.string(), 
    job_id : z.number(), 
    data : z.object({}).optional(), 

});

