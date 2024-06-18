    
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
// API
export interface APIItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    method : string ;
    source : object ;
    schema : object ;

}


// API
export const APISchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    method : z.string(), 
    source : z.object({}), 
    schema : z.object({}).optional(), 

});

