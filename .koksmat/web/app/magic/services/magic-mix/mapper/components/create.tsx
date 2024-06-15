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
    import {MapperForm} from "./form";
    
    import {MapperItem} from "../applogic/model";
    export default function CreateMapper() {
       
        const mapper = {} as MapperItem;
        return (
          <div>{mapper && 
          <MapperForm mapper={mapper} editmode="create"/>}
         
          </div>
        );
    }
