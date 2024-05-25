/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {DatasetForm} from "./form";

import {DatasetItem} from "../applogic/model";
export default function UpdateDataset(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<DatasetItem>(
      "magic-mix.dataset",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const dataset = readResult.data;
    return (
      <div>{dataset && 
      <DatasetForm dataset={dataset} editmode="update"/>}
     
      </div>
    );
}
