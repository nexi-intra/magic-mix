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
import {SQLForm} from "./form";

import {SQLItem} from "../applogic/model";
export default function UpdateSQL(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<SQLItem>(
      "magic-mix.sql",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const sql = readResult.data;
    return (
      <div>{sql && 
      <SQLForm sql={sql} editmode="update"/>}
     
      </div>
    );
}
