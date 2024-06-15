    
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
// User
export interface UserItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    system_id : number ;
    fullname : string ;

}


// User
export const UserSchema = z.object({  
   
        name : z.string(), 
    description : z.string(), 
    system_id : z.number(), 
    fullname : z.string(), 

});

