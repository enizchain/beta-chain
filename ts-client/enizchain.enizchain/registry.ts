import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePost } from "./types/invochain/invochain/tx";
import { MsgUpdatePost } from "./types/invochain/invochain/tx";
import { MsgDeletePost } from "./types/invochain/invochain/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/invochain.invochain.MsgCreatePost", MsgCreatePost],
    ["/invochain.invochain.MsgUpdatePost", MsgUpdatePost],
    ["/invochain.invochain.MsgDeletePost", MsgDeletePost],
    
];

export { msgTypes }