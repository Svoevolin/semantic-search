from fastapi import APIRouter, Depends, Request
from fastapi.responses import JSONResponse
from internal.delivery.http.dto.search import SearchRequest, SearchListResponse
from internal.domain.search import SearchServiceI

router = APIRouter()

def get_search_service(request: Request) -> SearchServiceI:
    return request.app.state.search_service

@router.post("/search", response_model=SearchListResponse)
async def search_handler(dto: SearchRequest, service: SearchServiceI = Depends(get_search_service)):
    result = await service.search(dto.model())

    body, headers = SearchListResponse.from_model(result)
    return JSONResponse(content=body.model_dump(), headers=headers)
