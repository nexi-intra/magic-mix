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
import {DatasetItem} from "../applogic/model";


/* yankiebar */

export default function ReadDataset(props: { id: number }) {
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
      <div>
          
    {dataset && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{dataset.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{dataset.description}</div>
    </div>    <div>
        <div className="font-bold" >connection</div>
        <div>{dataset.connection_id}</div>
    </div>    <div>
        <div className="font-bold" >transformer</div>
        <div>{dataset.transformer_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{dataset.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{dataset.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{dataset.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{dataset.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{dataset.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
