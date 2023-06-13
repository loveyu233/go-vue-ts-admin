import request from "../request.ts";
import {ResCarouselType} from "@/type/userApi.ts";

enum CarouselApi {
    GETALL = "/carousel/all"
}

export const GetCarouselAll = () => request.get<any, ResCarouselType>(CarouselApi.GETALL);