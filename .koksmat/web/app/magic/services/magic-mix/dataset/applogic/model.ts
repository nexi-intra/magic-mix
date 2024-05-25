    
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
// Dataset
export interface DatasetItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    connection_id : number ;
    transformer_id : number ;

}


// Dataset
export const DatasetSchema = z.object({  
   
        name : z.string(), 
    description : z.string(), 
    connection_id : z.number(), 
    transformer_id : z.number(), 

});

