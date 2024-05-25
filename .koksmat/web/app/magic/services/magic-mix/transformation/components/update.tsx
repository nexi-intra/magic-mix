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
import {TransformationForm} from "./form";

import {TransformationItem} from "../applogic/model";
export default function UpdateTransformation(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TransformationItem>(
      "magic-mix.transformation",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const transformation = readResult.data;
    return (
      <div>{transformation && 
      <TransformationForm transformation={transformation} editmode="update"/>}
     
      </div>
    );
}
