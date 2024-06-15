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
    import {TransformationForm} from "./form";
    
    import {TransformationItem} from "../applogic/model";
    export default function CreateTransformation() {
       
        const transformation = {} as TransformationItem;
        return (
          <div>{transformation && 
          <TransformationForm transformation={transformation} editmode="create"/>}
         
          </div>
        );
    }
