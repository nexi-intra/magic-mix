    
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
// Mapper
export interface MapperItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    source_id : number ;
    transformation_id : number ;
    target_id : number ;

}


// Mapper
export const MapperSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    source_id : z.number(), 
    transformation_id : z.number(), 
    target_id : z.number(), 

});

