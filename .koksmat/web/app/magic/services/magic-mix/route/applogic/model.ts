    
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
// Route
export interface RouteItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    method : string ;
    slug : string ;
    api_id : number ;

}


// Route
export const RouteSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    method : z.string(), 
    slug : z.string(), 
    api_id : z.number().optional(), 

});

