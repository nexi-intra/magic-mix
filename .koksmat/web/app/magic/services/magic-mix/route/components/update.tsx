/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {RouteForm} from "./form";

import {RouteItem} from "../applogic/model";
export default function UpdateRoute(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<RouteItem>(
      "magic-mix.route",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const route = readResult.data;
    return (
      <div>{route && 
      <RouteForm route={route} editmode="update"/>}
     
      </div>
    );
}
