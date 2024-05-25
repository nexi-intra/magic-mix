    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {ProcessLogItem} from "../applogic/model";


/* yankiebar */

export default function ReadProcessLog(props: { id: number }) {
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
      <div>
          
    {processlog && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{processlog.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{processlog.description}</div>
    </div>    <div>
        <div className="font-bold" >transformation</div>
        <div>{processlog.transformation_id}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{processlog.status}</div>
    </div>    <div>
        <div className="font-bold" >message</div>
        <div>{processlog.message}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{processlog.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{processlog.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{processlog.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{processlog.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{processlog.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
