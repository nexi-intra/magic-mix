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
    import {ColumnForm} from "./form";
    
    import {ColumnItem} from "../applogic/model";
    export default function CreateColumn() {
       
        const column = {} as ColumnItem;
        return (
          <div>{column && 
          <ColumnForm column={column} editmode="create"/>}
         
          </div>
        );
    }
