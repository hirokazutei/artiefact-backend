# -*- coding: utf-8 -*-
from sqlalchemy import Column, DateTime, ForeignKey
from sqlalchemy.dialects.postgresql import BIGINT, NUMERIC

from . import Base


class Saga(Base):
    "A Saga"

    __tablename__ = 'saga'

    id = Column(BIGINT, primary_key=True)
    user_id = Column(BIGINT, nullable=False, index=True)
    begin_date = Column(DateTime(timezone=True), nullable=False)
    end_date = Column(DateTime(timezone=True), nullable=True)
    starting_longitudes = Column(NUMERIC, nullable=False)
    starting_latitudes = Column(NUMERIC, nullable=False)
    ending_longitudes = Column(NUMERIC, nullable=True)
    ending_latitudes = Column(NUMERIC, nullable=True) 

class Chapter(Base):
    "A Chapter"

    __tablename__ = 'chapter'

    id = Column(BIGINT, primary_key=True)
    saga_id = Column(BIGINT, ForeignKey('saga.id'), nullable=False)
    begin_date = Column(DateTime(timezone=True), nullable=False) 
    end_date = Column(DateTime(timezone=True), nullable=True)
    starting_longitudes = Column(NUMERIC, nullable=False)
    starting_latitudes = Column(NUMERIC, nullable=False)
    ending_longitudes = Column(NUMERIC, nullable=True)
    ending_latitudes = Column(NUMERIC, nullable=True)