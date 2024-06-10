    
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
// ImportData
export interface ImportDataItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    data : object ;

}


// ImportData
export const ImportDataSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    data : z.object({}), 

});

