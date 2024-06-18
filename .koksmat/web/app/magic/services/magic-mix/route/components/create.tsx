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
    import {RouteForm} from "./form";
    
    import {RouteItem} from "../applogic/model";
    export default function CreateRoute() {
       
        const route = {} as RouteItem;
        return (
          <div>{route && 
          <RouteForm route={route} editmode="create"/>}
         
          </div>
        );
    }
