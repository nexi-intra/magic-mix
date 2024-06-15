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
import {SQLItem} from "../applogic/model";


/* yankiebar */

export default function ReadSQL(props: { id: number }) {
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
      <div>
          
    {sql && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{sql.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{sql.description}</div>
    </div>    <div>
        <div className="font-bold" >query</div>
        <div>{sql.query}</div>
    </div>                <div>
                    <div className="font-bold" >schema</div>
                    <div>{JSON.stringify(sql.schema,null,2)}</div>
                </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{sql.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{sql.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{sql.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{sql.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{sql.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
