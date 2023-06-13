function Pie() {
    let dataArr = [];
    for (let i = 0; i < 150; i++) {
        if (i % 2 === 0) {
            dataArr.push({
                name: (i + 1).toString(),
                value: 50,
                itemStyle: {
                    normal: {
                        color: "#00AFFF",
                        borderWidth: 0,
                        borderColor: "rgba(0,0,0,0)"
                    }
                }
            });
        } else {
            dataArr.push({
                name: (i + 1).toString(),
                value: 100,
                itemStyle: {
                    normal: {
                        color: "rgba(0,0,0,0)",
                        borderWidth: 0,
                        borderColor: "rgba(0,0,0,0)"
                    }
                }
            });
        }
    }
    return dataArr;
}

// 水球图
export const WaterBalloonDiagramOption = {
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    // backgroundColor: "#081736",
    // title: {
    //     text: "水球图",
    //     textStyle: {
    //         fontWeight: "normal",
    //         fontSize: 25,
    //         color: "rgb(97, 142, 205)"
    //     }
    // },
    series: [
        {
            // value: 50, //  内容 配合formatter
            type: "liquidFill",
            radius: "70%", // 控制中间圆球的尺寸（此处可以理解为距离外圈圆的距离控制）
            center: ["50%", "50%"],
            data: [
                0.8,
                {
                    value: 0.8,
                    direction: "left" //波浪方向
                }
            ], // data个数代表波浪数
            backgroundStyle: {
                borderWidth: 1,
                color: "rgba(62, 208, 255, 1)" // 球体本景色
            },
            amplitude: "6%", //波浪的振幅
            // 修改波浪颜色
            // color: ['#0286ea', 'l#0b99ff'], // 每个波浪不同颜色，颜色数组长度为对应的波浪个数
            color: [
                {
                    type: "linear",
                    x: 0,
                    y: 0,
                    x2: 0,
                    y2: 1,
                    colorStops: [
                        {
                            offset: 1,
                            color: "#6CDEFC"
                        },
                        {
                            offset: 0,
                            color: "#429BF7"
                        }
                    ],
                    globalCoord: false
                }
            ],
            label: {
                normal: {
                    // formatter: 0.87 * 100 + '\n%',
                    formatter: 0.8 * 100 + "\n{d|%}",
                    //  formatter: function(params){
                    //     return params.value* 100 + " \n%";
                    // },
                    rich: {
                        d: {
                            fontSize: 36
                        }
                    },
                    textStyle: {
                        fontSize: 48,
                        color: "#fff"
                    }
                }
            },
            outline: {
                show: false
            }
        },
        {
            type: "pie",
            z: 1,
            center: ["50%", "50%"],
            radius: ["72%", "73.5%"], // 控制外圈圆的粗细
            hoverAnimation: false,
            data: [
                {
                    name: "",
                    value: 500,
                    labelLine: {
                        show: false
                    },
                    itemStyle: {
                        color: "#00AFFF"
                    },
                    emphasis: {
                        labelLine: {
                            show: false
                        },
                        itemStyle: {
                            color: "#00AFFF"
                        }
                    }
                }
            ]
        },
        {
            //外发光
            type: "pie",
            z: 1,
            zlevel: -1,
            radius: ["70%", "90.5%"],
            center: ["50%", "50%"],
            hoverAnimation: false,
            clockWise: false,
            itemStyle: {
                normal: {
                    borderWidth: 20,
                    color: "rgba(224,242,255,0.3)"
                }
            },
            label: {
                show: false
            },
            data: [100]
        },
        {
            //底层外发光
            type: "pie",
            z: 1,
            zlevel: -2,
            radius: ["70%", "100%"],
            center: ["50%", "50%"],
            hoverAnimation: false,
            clockWise: false,
            itemStyle: {
                normal: {
                    borderWidth: 20,
                    color: "rgba(224,242,255,.2)"
                }
            },
            label: {
                show: false
            },
            data: [100]
        },
        {
            type: "pie",
            zlevel: 0,
            silent: true,
            radius: ["78%", "80%"],
            z: 1,
            label: {
                normal: {
                    show: false
                }
            },
            labelLine: {
                normal: {
                    show: false
                }
            },
            data: Pie()
        }
    ]
};


// 柱状比例图
export const HistogramOption = {
    // title: {
    //     text: "柱状比例图"
    // },
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    xAxis: {
        show: false,
        min: 0,
        max: 100
    },
    yAxis: {
        show: false,
        type: "category"
    },
    series: [
        {
            data: [60],
            type: "bar",
            barWidth: "30px",
            z: "100 ",
            itemStyle: {
                color: "skyblue",
                borderRadius: 20
            },
            label: {
                show: true
            }
        }, {
            data: [80],
            type: "bar",
            barWidth: "30px",
            barGap: "-100%",
            itemStyle: {
                color: "pink",
                borderRadius: 20
            },
            label: {
                show: true,
                offset: [100, 0]
            }
        }
    ],
    grid: {
        left: 20,
        right: 20,
        top: 0,
        bottom: 0
    }
};

// 饼图

export const PieChartOption = {
    legend: {
        top: "bottom"
    },
    // title: {
    //     text: "饼图"
    // },
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    series: [
        {
            name: "饼图",
            type: "pie",
            radius: [10, 100],
            center: ["50%", "50%"],
            roseType: "area",
            itemStyle: {
                borderRadius: 8
            },
            data: [
                {value: 40, name: "rose 1"},
                {value: 38, name: "rose 2"},
                {value: 32, name: "rose 3"},
                {value: 30, name: "rose 4"},
                {value: 28, name: "rose 5"},
                {value: 26, name: "rose 6"},
                {value: 22, name: "rose 7"},
                {value: 18, name: "rose 8"}
            ]
        }
    ],
    grid: {
        left: 0,
        right: 0,
        top: 0,
        bottom: 0
    }
};

// 飞机图标
let planePath = "path://M1705.06,1318.313v-89.254l-319.9-221.799l0.073-208.063c0.521-84.662-26.629-121.796-63.961-121.491c-37.332-0.305-64.482,36.829-63.961,121.491l0.073,208.063l-319.9,221.799v89.254l330.343-157.288l12.238,241.308l-134.449,92.931l0.531,42.034l175.125-42.917l175.125,42.917l0.531-42.034l-134.449-92.931l12.238-241.308L1705.06,1318.313z";

// 中国地图
export const ChinaMapOption = {
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    geo: {
        map: "china",
        roam: true,
        zoom: 1.6,
        top: 200,
        itemStyle: {
            color: "#de7a7a"
        },
        label: {
            show: true,
            fontSize: 15
        },
        emphasis: {
            label: {
                fontSize: 25,
                color: "red"
            }
        }
    },
    series: {
        type: "lines",
        data: [
            {
                coords: [
                    [116.405285, 39.904989],  // 起点
                    [91.132212, 29.660361]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: 0.5
                }
            },
            {
                coords: [
                    [126.642464, 45.756967],  // 起点
                    [113.665412, 34.757975]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: -0.3
                }
            },
            {
                coords: [
                    [113.665412, 34.757975],  // 起点
                    [91.132212, 29.660361]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: 0.5
                }
            }, {
                coords: [
                    [112.549248, 37.857014],  // 起点
                    [104.065735, 30.659462]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: 0.5
                }
            }, {
                coords: [
                    [117.190182, 39.125596],  // 起点
                    [121.472644, 31.231706]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: -0.5
                }
            }, {
                coords: [
                    [113.280637, 23.125178],  // 起点
                    [103.823557, 36.058039]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: 0.5
                }
            }, {
                coords: [
                    [87.617733, 43.792818],  // 起点
                    [126.642464, 45.756967]   // 终点
                ],
                // 统一的样式设置
                lineStyle: {
                    type: "dotted",
                    width: 0,
                    curveness: -0.5
                }
            }
        ],
        // 特效
        effect: {
            show: true,
            period: 4,
            color: "white",
            roundTrip: true,
            symbol: planePath,
            symbolSize: 40
        }
    }
};

export const LineChartOption = {
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    xAxis: {
        type: "category",
        data: ["星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"]
    },
    yAxis: {
        type: "value"
    },
    series: [
        {
            data: [820, 932, 901, 934, 1290, 2000, 2700],
            type: "line",
            smooth: true
        }, {
            data: [500, 1000, 901, 700, 1290, 800, 200],
            type: "line",
            smooth: true
        }, {
            data: [100, 2200, 701, 400, 990, 400, 2000],
            type: "line",
            smooth: true
        }, {
            data: [900, 500, 301, 900, 2090, 500, 1200],
            type: "line",
            smooth: true
        }
    ]
};


export const Box1Option = {
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    xAxis: {},
    yAxis: {},
    series: [
        {
            name: "Placeholder",
            type: "bar",
            data: [0, 1700, 1400, 1200, 300, 0]
        }
    ]
};


export const Box2Option = {
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    polar: {
        radius: [30, "80%"]
    },
    angleAxis: {
        max: 4,
        startAngle: 75
    },
    radiusAxis: {
        type: "category",
        data: ["a", "b", "c", "d"]
    },
    tooltip: {},
    series: {
        type: "bar",
        data: [2, 1.2, 2.4, 3.6],
        coordinateSystem: "polar",
        label: {
            show: true,
            position: "middle",
            formatter: "{b}: {c}"
        }
    }
};


const gaugeData = [
    {
        value: 20,
        name: "",
        title: {
            offsetCenter: ["0%", "-30%"]
        },
        detail: {
            valueAnimation: true,
            offsetCenter: ["0%", "-20%"]
        }
    },
    {
        value: 40,
        name: "",
        title: {
            offsetCenter: ["0%", "0%"]
        },
        detail: {
            valueAnimation: true,
            offsetCenter: ["0%", "10%"]
        }
    }
];
export const Box3Option = {
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    series: [
        {
            type: "gauge",
            startAngle: 90,
            endAngle: -270,
            pointer: {
                show: false
            },
            progress: {
                show: true,
                overlap: false,
                roundCap: true,
                clip: false,
                itemStyle: {
                    borderWidth: 1,
                    borderColor: "#464646"
                }
            },
            axisLine: {
                lineStyle: {
                    width: 40
                }
            },
            splitLine: {
                show: false,
                distance: 0,
                length: 10
            },
            axisTick: {
                show: false
            },
            axisLabel: {
                show: false,
                distance: 50
            },
            data: gaugeData,
            title: {
                fontSize: 10
            },
            detail: {
                width: 50,
                height: 14,
                fontSize: 20,
                color: "inherit",
                borderColor: "inherit",
                borderRadius: 20,
                borderWidth: 1,
                formatter: "{value}%"
            }
        }
    ],
    grid: {
        top: 200
    }
};


