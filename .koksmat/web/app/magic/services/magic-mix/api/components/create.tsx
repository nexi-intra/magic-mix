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
    import {APIForm} from "./form";
    
    import {APIItem} from "../applogic/model";
    export default function CreateAPI() {
       
        const api = {} as APIItem;
        return (
          <div>{api && 
          <APIForm api={api} editmode="create"/>}
         
          </div>
        );
    }
