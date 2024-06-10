    
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
// Transformation
export interface TransformationItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    input_id : number ;
    output_id : number ;

}


// Transformation
export const TransformationSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    input_id : z.number().optional(), 
    output_id : z.number().optional(), 

});

