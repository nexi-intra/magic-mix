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
import {APIForm} from "./form";

import {APIItem} from "../applogic/model";
export default function UpdateAPI(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<APIItem>(
      "magic-mix.api",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const api = readResult.data;
    return (
      <div>{api && 
      <APIForm api={api} editmode="update"/>}
     
      </div>
    );
}
