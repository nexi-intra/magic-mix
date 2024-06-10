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
import {ImportDataForm} from "./form";

import {ImportDataItem} from "../applogic/model";
export default function UpdateImportData(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ImportDataItem>(
      "magic-mix.importdata",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const importdata = readResult.data;
    return (
      <div>{importdata && 
      <ImportDataForm importdata={importdata} editmode="update"/>}
     
      </div>
    );
}
