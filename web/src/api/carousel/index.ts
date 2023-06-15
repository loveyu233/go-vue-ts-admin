import request from "../request.ts";
import {AddCarouselType, CarouselType, ResCarouselType} from "@/type/userApi.ts";

enum CarouselApi {
    GETALL = "/carousel/all",
    GETSHOW = "/carousel/show",
    ADD = "/carousel/add",
    DEL = "/carousel/del",
    UPDATE = "/carousel/update",
    OPERATE = "/carousel/delete"
}

export const GetCarouselAll = () => request.get<any, ResCarouselType>(CarouselApi.GETALL);
export const GetCarouselShow = () => request.get<any, ResCarouselType>(CarouselApi.GETSHOW);

export const AddCarousel = (data: CarouselType) => request.post<any, AddCarouselType>(CarouselApi.ADD, data);

export const DelCarousel = (id: number) => request.get<any, AddCarouselType>(CarouselApi.DEL, {params: {"id": id}});

export const UpdateCarousel = (data: CarouselType) => request.post<any>(CarouselApi.UPDATE, data);

export const OperateCarouselIn = (data: number[], operate: string) => request.get(CarouselApi.OPERATE, {
    params: {
        "ids": data,
        "operate": operate
    }
});