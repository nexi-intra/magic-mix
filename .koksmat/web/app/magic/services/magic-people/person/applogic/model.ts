    
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
// Person
export interface PersonItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    email : string ;
    displayname : string ;
    firstname : string ;
    lastname : string ;
    uniqueid : string ;
    nationality_id : number ;

}


// Person
export const PersonSchema = z.object({  
   
        name : z.string(), 
    description : z.string(), 
    email : z.string(), 
    displayname : z.string(), 
    firstname : z.string(), 
    lastname : z.string(), 
    uniqueid : z.string(), 
    nationality_id : z.number(), 

});

