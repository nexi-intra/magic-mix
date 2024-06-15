    
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
// Company
export interface CompanyItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    vatnumber : string ;
    phonenumber : string ;
    address : string ;
    city : string ;
    postalcode : string ;
    country_id : number ;

}


// Company
export const CompanySchema = z.object({  
   
        name : z.string(), 
    description : z.string(), 
    vatnumber : z.string(), 
    phonenumber : z.string(), 
    address : z.string(), 
    city : z.string(), 
    postalcode : z.string(), 
    country_id : z.number(), 

});

