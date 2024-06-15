    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {ColumnItem} from "../applogic/model";


/* yankiebar */

export default function ReadColumn(props: { id: number }) {
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
      <div>
          
    {column && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{column.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{column.description}</div>
    </div>    <div>
        <div className="font-bold" >datatype</div>
        <div>{column.datatype}</div>
    </div>    <div>
        <div className="font-bold" >sortorder</div>
        <div>{column.sortorder}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{column.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{column.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{column.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{column.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{column.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
