    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/app/koksmat/useservice";
    import { useState } from "react";
    import {SQLForm} from "./form";
    
    import {SQLItem} from "../applogic/model";
    export default function CreateSQL() {
       
        const sql = {} as SQLItem;
        return (
          <div>{sql && 
          <SQLForm sql={sql} editmode="create"/>}
         
          </div>
        );
    }
