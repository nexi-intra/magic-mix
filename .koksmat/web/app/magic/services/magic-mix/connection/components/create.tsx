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
    import {ConnectionForm} from "./form";
    
    import {ConnectionItem} from "../applogic/model";
    export default function CreateConnection() {
       
        const connection = {} as ConnectionItem;
        return (
          <div>{connection && 
          <ConnectionForm connection={connection} editmode="create"/>}
         
          </div>
        );
    }
