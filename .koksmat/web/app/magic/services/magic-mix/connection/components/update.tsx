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
import {ConnectionForm} from "./form";

import {ConnectionItem} from "../applogic/model";
export default function UpdateConnection(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ConnectionItem>(
      "magic-mix.connection",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const connection = readResult.data;
    return (
      <div>{connection && 
      <ConnectionForm connection={connection} editmode="update"/>}
     
      </div>
    );
}
