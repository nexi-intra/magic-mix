    
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
// Group
export interface GroupItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    hidden : boolean ;
    notesid : string ;

}


// Group
export const GroupSchema = z.object({  
   
        name : z.string(), 
    description : z.string(), 
    hidden : z.boolean(), 
    notesid : z.string(), 

});

