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
import {ProcessLogForm} from "./form";

import {ProcessLogItem} from "../applogic/model";
export default function UpdateProcessLog(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ProcessLogItem>(
      "magic-mix.processlog",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const processlog = readResult.data;
    return (
      <div>{processlog && 
      <ProcessLogForm processlog={processlog} editmode="update"/>}
     
      </div>
    );
}
