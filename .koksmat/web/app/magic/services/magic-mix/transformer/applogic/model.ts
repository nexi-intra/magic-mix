    
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
// Transformer
export interface TransformerItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    code : string ;

}


// Transformer
export const TransformerSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    code : z.string().optional(), 

});

