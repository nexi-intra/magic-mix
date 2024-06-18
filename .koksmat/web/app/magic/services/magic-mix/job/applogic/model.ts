    
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
// Job
export interface JobItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    status : string ;
    startat : string ;
    startedAt : string ;
    completedAt : string ;
    maxduration : number ;
    script : string ;
    data : object ;

}


// Job
export const JobSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    status : z.string(), 
    startat : z.string().optional(), 
    startedAt : z.string().optional(), 
    completedAt : z.string().optional(), 
    maxduration : z.number().optional(), 
    script : z.string().optional(), 
    data : z.object({}).optional(), 

});

