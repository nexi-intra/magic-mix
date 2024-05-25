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
import {TransformerForm} from "./form";

import {TransformerItem} from "../applogic/model";
export default function UpdateTransformer(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TransformerItem>(
      "magic-mix.transformer",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const transformer = readResult.data;
    return (
      <div>{transformer && 
      <TransformerForm transformer={transformer} editmode="update"/>}
     
      </div>
    );
}
