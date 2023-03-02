import { LNURL_SUCCESS, LNURL_FAILED} from '../actions/types'

const initialState = {
    lnurl: "",
    loading: true,
}

const lnurreducer =(state = initialState, action)=>{
    const { type, payload } = action;

    switch(type){
        case LNURL_SUCCESS:
            return {
                ...state,
                lnurl: payload.lnurl,
                loading: false
            }

        case LNURL_FAILED:
            default:
                return state
    }
}

export default lnurreducer;