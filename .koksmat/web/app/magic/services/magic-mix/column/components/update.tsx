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
import {ColumnForm} from "./form";

import {ColumnItem} from "../applogic/model";
export default function UpdateColumn(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ColumnItem>(
      "magic-mix.column",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const column = readResult.data;
    return (
      <div>{column && 
      <ColumnForm column={column} editmode="update"/>}
     
      </div>
    );
}
